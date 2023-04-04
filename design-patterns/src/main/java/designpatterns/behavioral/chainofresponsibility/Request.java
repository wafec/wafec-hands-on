package designpatterns.behavioral.chainofresponsibility;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Request {
	private String data;
}
