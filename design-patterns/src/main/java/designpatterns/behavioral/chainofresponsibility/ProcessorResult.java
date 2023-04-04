package designpatterns.behavioral.chainofresponsibility;

import lombok.Builder;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.experimental.Accessors;

@Getter
@EqualsAndHashCode
@Builder
public class ProcessorResult {
	@Accessors( fluent = true )
	private boolean shouldCallNextHandler;
}
