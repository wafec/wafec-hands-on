package designpatterns.behavioral.memento;

import static org.assertj.core.api.Assertions.assertThat;

import java.time.Clock;
import java.time.Instant;
import java.time.ZoneId;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class CaretakerTest {
	private Caretaker testObj;

	private static final String TEST_DATA_1 = "test_data_1";
	private static final String TEST_DATA_2 = "test_data_2";
	private static final String NAME_1 = "name_1";
	private static final String NAME_2 = "name_2";
	private static final Clock CLOCK_1 = Clock.fixed( Instant.parse( "2018-01-01T10:00:00Z" ), ZoneId.systemDefault() );
	private static final Clock CLOCK_2 = Clock.fixed( Instant.parse( "2019-01-01T10:00:00Z" ), ZoneId.systemDefault() );

	@BeforeEach
	public void setUp() {
		testObj = new Caretaker();
	}

	@Test
	void addSnapshots() {
		final DataObject o1 = getDataObject1();
		final DataObject o2 = getDataObject2();

		testObj.add( o2.createSnapshot() );
		testObj.add( o1.createSnapshot() );

		assertThat( testObj.first() ).isEqualTo( o1.createSnapshot() );
		assertThat( testObj.last() ).isEqualTo( o2.createSnapshot() );
	}

	@Test
	void repeatedElements() {
		final DataObject o1 = getDataObject1();

		testObj.add( o1.createSnapshot() );
		testObj.add( o1.createSnapshot() );

		assertThat( testObj.size() ).isEqualTo( 1 );
	}

	private static DataObject getDataObject1() {
		return DataObject.builder()
				.data( TEST_DATA_1 )
				.clock( CLOCK_1 )
				.name( NAME_1 )
				.build();
	}

	private static DataObject getDataObject2() {
		return DataObject.builder()
				.data( TEST_DATA_2 )
				.clock( CLOCK_2 )
				.name( NAME_2 )
				.build();
	}
}
