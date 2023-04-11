package designpatterns.behavioral.strategy;

class StrategyClient {
	private Strategy strategy;

	public static final String DATA_TO_SEND = "The strategy pattern rocks!";

	public StrategyClient( final Strategy strategy ) {
		this.strategy = strategy;
	}

	public void doWork() {
		strategy.performThisStrategy( StrategyData.builder().data( DATA_TO_SEND ).build() );
	}

	public void changeStrategy( final Strategy strategy ) {
		this.strategy = strategy;
	}
}
