package designpatterns.behavioral.visitor;

import lombok.Builder;
import lombok.Getter;

@Getter
@Builder
public class VisitedB {
	private String data;

	public void accept( final Visitor visitor ) {
		visitor.doForB( this );
	}
}
