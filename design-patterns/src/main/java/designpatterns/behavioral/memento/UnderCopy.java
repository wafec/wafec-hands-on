package designpatterns.behavioral.memento;

public interface UnderCopy {
	Memento createSnapshot();
	void restore( Memento memento );
}
