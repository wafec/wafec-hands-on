package designpatterns.behavioral.memento;

import java.sql.Date;
import java.time.Clock;

import lombok.Builder;
import lombok.EqualsAndHashCode;
import lombok.Getter;

@Getter
@Builder
@EqualsAndHashCode
class DataObject implements UnderCopy {
	private String data;
	private Clock clock;
	private String name;

	@Override
	public Memento createSnapshot() {
		return Snapshot.builder()
				.data( data )
				.createdDate( Date.from( clock.instant() ) )
				.name( name )
				.build();
	}

	@Override
	public void restore( final Memento memento ) {
		final Snapshot snapshot = (Snapshot) memento;
		this.data = snapshot.getData();
	}
}
