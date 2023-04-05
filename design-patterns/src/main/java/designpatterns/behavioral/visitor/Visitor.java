package designpatterns.behavioral.visitor;

/*
	Main idea: a visitor class to have access to visited nodes to perform complementary operation
	without having to change existent classes.
	Pros:
		1. Open/Closed Principle: once the classes have settled down the signature for accepting the
		 						  visitors all the remaining interaction won't require any further
		 						  change to existing classes
		2. Single Responsibility: it allows to have multiple classes very well scoped
		3. Traverse complex structure: the visitor might traverse a complex structure and obtaining
									   information from all objects to compose a new thing
 */
public interface Visitor {
	void doForA( VisitedA visited );
	void doForB( VisitedB visited );
}
