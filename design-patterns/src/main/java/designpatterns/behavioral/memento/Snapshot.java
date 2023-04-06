package designpatterns.behavioral.memento;

import java.util.Date;

import lombok.Builder;
import lombok.EqualsAndHashCode;
import lombok.Getter;

@Getter
@Builder
@EqualsAndHashCode
public class Snapshot implements Memento {
	private String data;
	private String name;
	private Date createdDate;

	@Override
	public String getName() {
		return name;
	}

	@Override
	public Date getSnapshotDate() {
		return createdDate;
	}
}
