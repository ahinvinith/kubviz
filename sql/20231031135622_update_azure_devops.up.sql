ALTER TABLE azure_devops
    ENGINE = MergeTree()
    ORDER BY (ClusterName, TimeStamp)
    TTL TimeStamp + INTERVAL 30 DAY
    SETTINGS index_granularity = 8192;