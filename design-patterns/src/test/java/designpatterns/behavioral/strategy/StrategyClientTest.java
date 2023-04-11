package designpatterns.behavioral.strategy;

import static org.mockito.Mockito.verify;

import java.util.function.Consumer;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

@ExtendWith( MockitoExtension.class )
public class StrategyClientTest {
	private StrategyClient testObj;

	@Mock
	private Strategy strategy;

	@BeforeEach
	public void setUp() {
		testObj = new StrategyClient( strategy );
	}

	@Test
	void callingStrategy() {
		final StrategyData expectedData = StrategyData.builder().data( StrategyClient.DATA_TO_SEND ).build();

		testObj.doWork();

		verify( strategy ).performThisStrategy( expectedData );
	}

	@Nested
	class ConcreteStrategiesTest {
		private StrategyA strategyA;
		private StrategyB strategyB;

		@Mock
		private Consumer<String> consumerA;
		@Mock
		private Consumer<String> consumerB;

		@BeforeEach
		public void setUp() {
			strategyA = new StrategyA( consumerA );
			strategyB = new StrategyB( consumerB );
		}

		@Test
		void switchingBetweenStrategies() {
			final String expectedResult = StrategyClient.DATA_TO_SEND;

			testObj.changeStrategy( strategyA );
			testObj.doWork();
			testObj.changeStrategy( strategyB );
			testObj.doWork();

			verify( consumerA ).accept( expectedResult );
			verify( consumerB ).accept( expectedResult );
		}
	}
}
