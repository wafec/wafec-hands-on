package designpatterns.behavioral.iterator;

import static org.assertj.core.api.Assertions.assertThat;

import java.util.List;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class ListInOrderIteratorTest {
	private ListInOrderIterator<Integer> testObj;
	private static final Integer FIRST_ELEMENT = 1;
	private static final Integer SECOND_ELEMENT = 2;
	private static final Integer THIRD_ELEMENT = 3;
	private static final List<Integer> LIST = List.of( FIRST_ELEMENT, SECOND_ELEMENT, THIRD_ELEMENT );

	@BeforeEach
	public void setUp() {
		testObj = new ListInOrderIterator<>( LIST );
	}

	@Test
	void traversing() {
		for ( int i = 0; i < 2; i++ ) {
			final Integer value = testObj.get();
			final boolean hasNext = testObj.next();
			assertThat( value ).isEqualTo( LIST.get( i ) );
			assertThat( hasNext ).isTrue();
		}
		final Integer value = testObj.get();
		final boolean hasNext = testObj.next();
		assertThat( value ).isEqualTo( THIRD_ELEMENT );
		assertThat( hasNext ).isFalse();
	}
}
