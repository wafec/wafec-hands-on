package designpatterns.behavioral.memento;

import java.util.Date;

/*
	Main idea: allows to copy and restore private data to external parties without revealing
			   details of implementation
	Pros:
		1. No violation to encapsulation: snapshots are created without going into the object
										  for getting the data
		2. The origin does not need to care about keeping the history of its snapshots
	Cons:
		1. Lots of RAM used
		2. Caretaker should track the origin to destroy obsolete objects
 */
public interface Memento {
	String getName();
	Date getSnapshotDate();
}
