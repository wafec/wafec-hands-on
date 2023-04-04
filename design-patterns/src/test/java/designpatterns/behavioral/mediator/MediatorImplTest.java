package designpatterns.behavioral.mediator;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.when;

import java.util.List;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

@ExtendWith( MockitoExtension.class )
public class MediatorImplTest {
	private MediatorImpl testObj;

	@Mock
	private ComponentA componentA;

	@Mock
	private ComponentB componentB;

	private static final int VALUE_OF_A = 10;

	private static final int VALUE_OF_B = 20;

	@BeforeEach
	public void setUp() {
		testObj = new MediatorImpl( componentA, componentB );
	}

	@Test
	void getListOfValuesInCorrectOrder() {
		setUpComponentA();
		setUpComponentB();
		final List<Integer> expectedValuesList = List.of( VALUE_OF_A, VALUE_OF_B );

		testObj.notifyMediator( componentA );
		testObj.notifyMediator( componentB );
		final List<Integer> valuesList = testObj.getValuesList();

		assertThat( valuesList ).isEqualTo( expectedValuesList );
	}

	private void setUpComponentA() {
		when( componentA.getValue() ).thenReturn( VALUE_OF_A );
	}

	private void setUpComponentB() {
		when( componentB.getValue() ).thenReturn( VALUE_OF_B );
	}
}
