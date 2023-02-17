from dataset_util.utils_mysql import *
from utils.utils_dict import order_dumps, extract_special_config

result_dt_special_keys = [
    'id',
    'dataset_name',
    'instance',
    'run_id',
    'batch_run_id',
    'create_time',
    'max_depth',
    'tag',
    'seed',
    'score',
]


def construct_result_dt_dataframe():
    df = pd.DataFrame(columns=[
		'id',
		'dataset_name',
		'instance',
		'run_id',
		'batch_run_id',
		'create_time',
		'max_depth',
		'tag',
		'seed',
		'score',
    ])

    return df

def construct_dataframe_row(conf: dict):

    """
    构造 mysql 插入数据的结构
    """

    df = construct_result_dt_dataframe()

    row = extract_special_config(conf, result_dt_special_keys)

    df = df.append(row, ignore_index=True)

    return df


def insert(conf: dict):
    """
    插入一条新的数据
    """
    engine = get_mysql_engine('bse_bne')
    df = construct_dataframe_row(conf)
    df.to_sql('result_dt', engine, if_exists='append', index=False)


def query(conf: dict):
    """
    查询
    """
    engine = get_mysql_engine('bse_bne')

    sql = "select * from result_dt"

    if len(conf) > 0:
        sql += " where "
        conditions = []
        for k, v in conf.items():
            conditions.append(f"{k} = '{v}'")

        sql += ' AND '.join(conditions)

    sql += " order by create_time desc"

    df = pd.read_sql(sql, engine)
    # print(df)
    return df