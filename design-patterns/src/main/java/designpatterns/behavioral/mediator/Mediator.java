package designpatterns.behavioral.mediator;

/*
	Main idea: decouple objects that do not need to know each other
	Pros:
		1. Loose Coupling: components are not linked to each other
		2. Open/Closed Principle: new mediators might be added without changing any component
		3. Single Responsibility: objects have a well-defined scope
		4. Reuse: components are easy to reuse
	Cons:
		1. Can evolve into a God Object
 */
public interface Mediator {
	void notifyMediator( Component sender );
}
