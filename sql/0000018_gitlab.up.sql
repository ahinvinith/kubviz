CREATE TABLE IF NOT EXISTS gitlab (
	Author       String,
	Provider     String,
	CommitID     String,
	CommitUrl    String,
	EventType    String,
	RepoName     String,
	TimeStamp    DateTime('UTC'),
	Event        String,
	ExpiryDate DateTime DEFAULT now() + INTERVAL {{.TTLValue}} {{.TTLUnit}}
) ENGINE = MergeTree() 
ORDER BY ExpiryDate 
TTL ExpiryDate;