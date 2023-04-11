package designpatterns.behavioral.strategy;

import java.util.function.Consumer;

class StrategyA implements Strategy {
	private final Consumer<String> consumer;

	public StrategyA( final Consumer<String> consumer ) {
		this.consumer = consumer;
	}

	@Override
	public void performThisStrategy( final StrategyData strategyData ) {
		consumer.accept( strategyData.getData() );
	}
}
