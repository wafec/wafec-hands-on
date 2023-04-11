package designpatterns.behavioral.visitor;

import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

class JsonReporterVisitor implements Visitor {
	private final Map<String, String> fieldsByKey;

	public JsonReporterVisitor() {
		fieldsByKey = new HashMap<>();
	}

	@Override
	public void doForA( final VisitedA visited ) {
		fieldsByKey.put( "objectA", visited.getData() );
	}

	@Override
	public void doForB( final VisitedB visited ) {
		fieldsByKey.put( "objectB", visited.getData() );
	}

	public String getJsonString() {
		return "{" + fieldsByKey.entrySet().stream()
				.sorted( Map.Entry.comparingByKey() )
				.map( e1 -> String.format( "\"%s\": \"%s\"", e1.getKey(), e1.getValue() ) )
				.collect( Collectors.joining( ", " ) ) + "}";
	}
}
