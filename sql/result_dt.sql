CREATE TABLE result_dt
(
    id           bigint       NOT NULL AUTO_INCREMENT COMMENT '主键',
    run_id       varchar(64)  NOT NULL DEFAULT '' COMMENT '运行id',
    batch_run_id varchar(64)  NOT NULL DEFAULT '' COMMENT '批量运行id',
    seed         int(10) NOT NULL DEFAULT 0 COMMENT '随机种子',
    create_time  timestamp    NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    tag          varchar(10)  NOT NULL DEFAULT '' COMMENT '标记',
    dataset_name varchar(255) NOT NULL DEFAULT '' COMMENT '数据集名称',
    instance     int          NOT NULL DEFAULT 0 COMMENT '记录数',
    dim          int          NOT NULL DEFAULT 0 COMMENT '维度数',
    max_depth    int          NOT NULL DEFAULT 0 COMMENT 'max_depth',
    score        double       NOT NULL DEFAULT 0 COMMENT 'k',
    PRIMARY KEY (id)
) ENGINE=innoDB DEFAULT CHARSET=utf8 comment 'dt运行结果保存';
