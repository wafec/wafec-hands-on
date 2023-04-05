package designpatterns.behavioral.visitor;

import lombok.Builder;
import lombok.Getter;

@Getter
@Builder
public class VisitedA {
	private String data;

	public void accept( final Visitor visitor ) {
		visitor.doForA( this );
	}
}
