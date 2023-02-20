package tmpl

const BaseString = `
import pymysql
import pandas as pd
from sqlalchemy import create_engine
from model._tool import *

Default_DB_Name = "rht"

class Base:

   def __init__(self, db_name: str, table_name: str):
       self._db_name = db_name
       self._table_name = table_name

   def query(self, db_where: dict, order_by='id', columns=list) -> pd.DataFrame:
       sql = construct_query_sql(self._table_name, columns, db_where, order_by)
       df = pd.read_sql(sql, get_mysql_engine(self._db_name))
       return df

   def update(self, db_set: dict, db_where: dict):
       sql = construct_update_sql(self._table_name, db_set, db_where)
       execute(self.get_db_info(), sql)

   def insert(self, conf):
       sql = construct_insert_sql(self._table_name, conf)
       execute(self.get_db_info(), sql)

   def get_db_info(self):
       return get_db_info(self._db_name)


def get_db_info(database) -> dict:

   """
   获取数据库配置
   """

   db_info = {
       'host': '192.168.31.23',
       'user': 'root',
       'port': 3306,
       'password': '123456',
       'database': database,
       'charset': 'utf8',
   }

   return db_info


def get_mysql_engine(database):
   db_info = get_db_info(database)

   engine = create_engine(
       'mysql+pymysql://%(user)s:%(password)s@%(host)s:%(port)d/%(database)s?charset=utf8' % db_info,
       encoding='utf-8'
   )

   return engine


def execute(db_info: dict, sql: str):
   conn = pymysql.connect(**db_info)
   cursor = conn.cursor()
   cursor.execute(sql)
   conn.commit()
   cursor.close()
   conn.close()

`

const ToolString = `
def construct_db_where(db_where: dict) -> str:

    if len(db_where) == 0:
        return ""

    where_str = "where "

    conditions = []
    for k, v in db_where.items():
        conditions.append(f"{k} = '{v}'")

    where_str += ' AND '.join(conditions)

    return where_str


def construct_insert_sql(table_name: str, conf: dict) -> str:
    if len(conf) == 0:
        return ""

    keys = []
    values = []

    for k, v in conf.items():
        keys.append(k)
        values.append(f"'{v}'")

    keys_str = ','.join(keys)
    values_str = ','.join(values)
    insert_sql = f'insert into {table_name}({keys_str}) values({values_str})'
    return insert_sql


def construct_db_set(db_set) -> str:

    if len(db_set) == 0:
        return ""

    set_str = "set "

    conditions = []
    for k, v in db_set.items():
        conditions.append(f"{k} = '{v}'")

    set_str += ','.join(conditions)

    return set_str


def construct_update_sql(table_name: str, db_set: dict, db_where: dict) -> str:
    sql = f"update {table_name} "
    sql += construct_db_set(db_set) + " "
    sql += construct_db_where(db_where)
    return sql


def construct_query_sql(table_name: str, columns=list, db_where=None, order_by='id') -> str:
    if db_where is None:
        db_where = {}

    keys = '*'
    if type(columns) == list:
        keys = ','.join(columns)

    sql = f"select {keys} from {table_name} "
    sql += construct_db_where(db_where) + " "
    sql += f"order by {order_by}"

    return sql
`

const PropertiesString = `
Table_Name = "{{.DDL.NewName.Name}}"

Columns_Name = [
    {{- range .Columns}}
    '{{.Name}}',
    {{- end}}
]
`

const TString = `
import pandas as pd
from model.{{.DDL.NewName.Name}}.properties import *
from model._base import *


class {{s2t .DDL.NewName.Name}}(Base):

    def __init__(self, db_name: str):
        super().__init__(db_name, Table_Name)

    def query_by_run_id(self, run_id: str) -> pd.DataFrame:
        db_where = {
            'run_id': run_id,
        }
        df = self.query(db_where)
        return df



__{{.DDL.NewName.Name}} = None


def get_{{.DDL.NewName.Name}}() -> {{.DDL.NewName.Name}}:
    global __{{.DDL.NewName.Name}}

    if __{{.DDL.NewName.Name}} is None:
        __{{.DDL.NewName.Name}} = {{.DDL.NewName.Name}}(Default_DB_Name)
        return __{{.DDL.NewName.Name}}

    return __{{.DDL.NewName.Name}}

`
