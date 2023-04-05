package designpatterns.behavioral.visitor;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class JsonReporterVisitorTest {
	private JsonReporterVisitor testObj;

	private static final String TEST_DATA_FOR_A = "data_for_a";
	private static final String TEST_DATA_FOR_B = "data_for_b";

	@BeforeEach
	public void setUp() {
		testObj = new JsonReporterVisitor();
	}

	@Test
	void visitingAll() {
		final VisitedA visitedA = VisitedA.builder().data( TEST_DATA_FOR_A ).build();
		final VisitedB visitedB = VisitedB.builder().data( TEST_DATA_FOR_B ).build();
		final String expectedJsonString = "{\"objectA\": \"data_for_a\", \"objectB\": \"data_for_b\"}";

		visitedA.accept( testObj );
		visitedB.accept( testObj );
		final String jsonString = testObj.getJsonString();

		assertThat( jsonString ).isEqualTo( expectedJsonString );
	}
}
