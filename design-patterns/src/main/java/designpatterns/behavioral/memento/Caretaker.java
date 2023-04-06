package designpatterns.behavioral.memento;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

public class Caretaker {
	private final List<Memento> snapshotList;

	public Caretaker() {
		snapshotList = new ArrayList<>();
	}

	public void add( final Memento snapshot ) {
		if ( snapshotList.contains( snapshot ) )
			return;
		final int indexOf = Collections.binarySearch( snapshotList, snapshot, Comparator.comparing( Memento::getSnapshotDate ) );
		snapshotList.add( Math.abs( indexOf + 1 ), snapshot );
	}

	public Memento get( final int index ) {
		return snapshotList.get( index );
	}

	public int size() {
		return snapshotList.size();
	}

	public Memento last() {
		return get( size() - 1 );
	}

	public Memento first() {
		return get( 0 );
	}
}
