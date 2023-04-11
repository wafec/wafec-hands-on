package designpatterns.behavioral.strategy;

/*
	Main idea: multiple implementations of an algorithm can come separate as a plug and play piece
	Pros:
		1. Switch on and off different implementations
		2. Open/Closed Principle: new algorithms can come with no changes in the clients
		3. Favor composition
		4. Implementation details is isolated from who uses the algorithms
	Cons:
		1. Clients must know the algorithms and what they do
		2. Creating multiple classes for something simple might be overkill
		3. Considerably boilerplate code bloating the code
 */
public interface Strategy {
	void performThisStrategy( StrategyData strategyData );
}
