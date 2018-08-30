/*矿池信息*/
create table pools (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  UDID       varchar(64)  unique,--设备id
  name       varchar(64) , --矿池名称
	rate       INTEGER , --管理费
	feeAccount varchar(255) , --管理费账户 
	account    varchar(255) , --账户   
	password   varchar(255) ,--密码
  grade      varchar(64) , --评级级别
  created_at timestamp     --创建日期   
);
/* 矿机矿池关联表*/
create table miner (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  UDID       varchar(64)  , --矿机设备id
  pool_id    INTEGER,  --矿池id
	account    varchar(255),    --系统收益账户
  remaining_time   integer,  --剩余挖矿时间
  update_time timestamp, --更新写入时间
  status   bool , --是否挖矿 状态 1 是 0否
  in_time timestamp   , --进入矿池时间
  out_time timestamp  ,--退出矿池时间
  computing_power  INTEGER, --算力
  has_join bool , --是否已加入矿池
  exit_status integer -- 退出矿池状态 0空 1申请 2退出 默认0
);
 /*挖矿日志*/
create table miningLog (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  UDID       varchar(64)  , --矿机设备id
  pool_id    INTEGER, --矿池id
	start_time timestamp,    --开始挖矿时间
  stop_time  timestamp,  --停止挖矿时间
	mining_time  INTEGER  --挖矿时长
);
/*挖矿记录*/
create table poolMining (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  pool_id    INTEGER, --矿池id
  UDID       varchar(64), --矿机设备id
  block_number INTEGER, --挖矿区块号
  income     double, --挖到收入 
  date       timestamp  --挖到时间
);
/*收益分配记录*/
create table allocateIncome (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  mining_id INTEGER,  --挖到日志ID
  UDID       varchar(64)  , --矿机设备id
  pool_id    INTEGER, --矿池id
  income     double,  --分到收益
  fee        double,  --管理费
  --feeAccount varchar(255) , --管理费账户 
  --	account    varchar(255) , --账户  
  status     INTEGER,   --状态：1已结算 2未结算
  date       timestamp   --分配时间
); 