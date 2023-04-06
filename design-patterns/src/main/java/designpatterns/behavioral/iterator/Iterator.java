package designpatterns.behavioral.iterator;

/*
	Main idea: delegates to an external object the traversal of a collection structure
	Pros:
		1. Single Responsibility: separate the efficient data storage from the traversal algorithm
		2. Open/Closed Principle: new iterators for a collection might not impact the collection
		3. Parallel traversal: each iterator keeps its own state
		4. Stop and play anytime
	Cons:
		1. Overkill for simple collections that do not change too often
		2. Might be less efficient than just traversing using the original storage collection method
 */
public interface Iterator<T> {
	T get();
	boolean next();
}
