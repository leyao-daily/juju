(death-and-destruction)=
# Death and destruction

This document describes in detail the operations associated with the destruction
and removal of the fundamental state entities, and what agents are responsible
for those operations.

Each entity has an associated remove-* command. The precise implications of
removal differ by entity, but there are common features:

  * Only Alive entities can be removed; if removal is already in progress,
    as evidenced by an entity not being Alive, its "removal" is a no-op.
  * Entities might be removed immediately when they are destroyed, but this is not
    guaranteed.
  * If an entity is not removed immediately when it is destroyed, its eventual
    removal is very likely; but it is not currently guaranteed, for the
    following reasons:
      * Hardware failure, even when detected and corrected by a Provisioner, can
        lead to unremovable relations, because the redeployed unit doesn't know
        what relations it's in. This would be fixable by making the unit agent
        always leave the scope of relations when they're detected; or, probably
        better, by using actual remote scope membership state to track relation
        membership (rather than using the existence of a local directory, whose
        true intent is to track the membership of *other* units, as a proxy).
        This is actually a pretty serious BUG and should be addressed soon;
        neither proposed solution is very challenging.
      * Undetected hardware failure is annoying, and can block progress at any
        time, but can be observed via additional monitoring and resolved via out-
        of-band termination of borked resources, which should be sufficient to
        get the system moving again (assuming the above bug is fixed).
      * Unknown problems in juju, in which agents fail to fulfil the duties laid
        out in this document, could block progress at any time. Assuming a
        version of the agent code which does not exhibit the problem exists, it
        should always be possible to work around this situation by upgrading the
        agent; and, if that fails, by terminating the underlying provider
        resources out-of-band, as above, and waiting for the new agent version
        to be deployed on a fresh system (with the same caveat as above).
      * In light of the preceding two points, we don't *have* to implement
        "--force" options for `juju remove-machine` and `juju remove-unit`.
        This is good, because it will be tricky to implement them well.

In general, the user can just forget about entities once she's destroyed them;
the only caveat is that she may not create new applications with the same name, or
new relations identical to the destroyed ones, until those entities have
finally been removed.

In rough order of complexity, here's what happens when each entity kind is
destroyed. Note that in general the appropriate action is contingent on
mutable remote state, and many operations must be expressed as a transaction
involving several documents: the state API must be prepared to handle aborted
transactions and either diagnose definite failure or retry until the operation
succeeds (or, perhaps, finally error out pleading excessive contention).


## `juju remove-machine`


Removing a machine involves a single transaction defined as follows:

  * If the machine is not Alive, abort without error.
  * If the machine is the last one with JobManageModel, or has any assigned
    units, abort with an appropriate error.
  * Set the machine to Dying.

When a machine becomes Dying, the following operation occurs:

  * The machine's agent sets the machine to Dead.

When a machine becomes Dead, the following operations occur:

  * The machine's agent terminates itself and refuses to run again.
  * A Provisioner (a task running in some other machine agent) observes the
    death, decommissions the machine's resources, and removes the machine.

Removing a machine involves a single transaction defined as follows:

  * If the machine is not Dead, abort with an appropriate error.
  * Delete the machine document.


## `juju remove-unit`

Removing a unit involves a single transaction defined as follows:

  * If the unit is not Alive, abort without error.
  * Set the unit to Dying.

When a unit becomes Dying, the following operations occur:

  * The unit's agent leaves the scopes of all its relations. Note that this is
    a potentially complex sequence of operations and may take some time; in
    particular, any hooks that fail while the unit is leaving relations and
    stopping the charm will suspend this sequence until resolved (just like
    when the unit is Alive).
  * The unit's agent then sets the unit to Dead.

When a unit becomes Dead, the following operations occur:

  * The unit's agent terminates itself and refuses to run again.
  * The agent of the entity that deployed the unit (that is: a machine agent,
    for a principal unit; or, for a subordinate unit, the agent of a principal
    unit) observes the death, recalls the unit, and removes it.

Removing a unit involves a single transaction, defined as follows:

  * If the unit is a principal unit, unassign the unit from its machine.
  * If the unit is a subordinate unit, unassign it from its principal unit.
  * Delete the unit document.
  * If its application is Alive, or has at least two units, or is in at least
    one relation, decrement the application's unit count; otherwise remove the
    application.


## `juju remove-relation`

Removing a relation involves a single transaction defined as follows:

  * If the relation is not Alive, abort without error.
  * If any unit is in scope, set the relation to Dying.
  * Otherwise:
      * If the relation destruction is a direct user request, decrement the
        relation counts of both applications.
      * If the relation destruction is an immediate consequence of application
        destruction, decrement the reference count of the counterpart application
        alone. (This is because the application destruction logic is responsible
        for the relation count of the application being destroyed.)
      * Delete the relation document.
      * Mark the relation's unit settings documents for future cleanup.
          * This is done by creating a single document for the attention of
            some other part of the system (BUG: which doesn't exist), that is
            then responsible for mass-deleting the (potentially large number
            of) settings documents. This completely bypasses the mgo/txn
            mechanism, but we don't care because those documents are guaranteed
            to be unreferenced and unwatched, by virtue of the relation's prior
            removal.

When a relation is set to Dying, the following operations occur:

  * Every unit agent whose unit has entered the scope of that relation
    observes the change and causes its unit to leave scope.
  * If the relation has container scope, and no other container-scoped relation
    between its applications is Alive, the unit agents of the subordinate units in
    the relation will observe the change and destroy their units.

The Dying relation's document is finally removed in the same transaction in
which the last unit leaves its scope. Because this situation involves the
relation already being Dying, its applications may also be Dying, and so the
operations involved are subtly different to those taken above (when we know
for sure that the relation -- and hence both applications -- are still Alive).

  * Here, "the application" refers to the application of the unit departing scope, and
    "the counterpart application" refers to the other application in the relation.
  * Decrement the relation count of the unit's application (we know that application
    is not ready to be removed, because its unit is responsible for this
    transaction and the application clearly therefore has a unit count greater
    than zero).
  * Delete the relation document.
  * Mark the relation's unit settings documents for future cleanup.
  * If the counterpart application (the one that is not the unit's application) is
    Alive, or has at least one unit, or is in at least two relations, decrement
    its relation count; otherwise remove the counterpart application.


## `juju remove-application`

Removing an application involves a single transaction defined as follows:

  * If the application is not alive, abort without error.
  * If the application is in any relations, do the following for each one:
      * If the relation is already Dying, skip it.
      * If the relation is Alive, destroy the relation without modifying the
        application's relation count. If the relation's destruction implies its
        removal, increment a local removed-relations counter instead.
  * If the application's unit count is greater than 0, or if the value of the
    aforementioned removal counter is less than the application's relation count,
    we know that some entity will still hold a reference to the application after
    the transaction completes, so we set the application to Dying and decrement
    its relation count by the value of the removal counter.
  * Otherwise, remove the application immediately, because we know that no
    reference to the application will survive the transaction.

When an application becomes Dying, the following operations occur:

  * Every unit agent of the application observes the change and destroys its unit.

The Dying application's document is finally removed in the same transaction that
removes the last entity referencing that application. This could be either the
removal of the last unit in the application, or the removal of the last relation
the application is in, as described above. To remove an application, the following
operations must occur in a single transaction:

  * Remove the application document.
  * Remove the application's settings document.
