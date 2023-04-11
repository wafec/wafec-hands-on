package designpatterns.behavioral.iterator;

import java.util.List;

class ListInOrderIterator<T> implements Iterator<T> {
	private final List<T> list;
	private int index;

	public ListInOrderIterator( final List<T> list ) {
		this.list = List.copyOf( list );
		index = 0;
	}

	@Override
	public T get() {
		return list.get( index );
	}

	@Override
	public boolean next() {
		return ++index < list.size();
	}
}
