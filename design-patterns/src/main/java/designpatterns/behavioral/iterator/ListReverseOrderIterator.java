package designpatterns.behavioral.iterator;

import java.util.List;

public class ListReverseOrderIterator<T> implements Iterator<T> {
	private final List<T> list;
	private int index;

	public ListReverseOrderIterator( final List<T> list ) {
		this.list = list;
		index = list.size() - 1;
	}

	@Override
	public T get() {
		return list.get( index );
	}

	@Override
	public boolean next() {
		return --index >= 0;
	}
}
