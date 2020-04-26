CREATE TABLE IF NOT EXISTS tasks (contest VARCHAR(64), task TINYINT, title TINYTEXT, problem TEXT, time_limit INT, accuracy TINYINT, testcases JSON, PRIMARY KEY(contest, task)) DEFAULT CHARACTER SET utf8mb4;
CREATE TABLE IF NOT EXISTS submissions (id VARCHAR(64) PRIMARY KEY, contest VARCHAR(64), task TINYINT, lang VARCHAR(32), code TEXT, whole_result TINYINT, max_memory INT, max_time INT) DEFAULT CHARACTER SET utf8mb4;
