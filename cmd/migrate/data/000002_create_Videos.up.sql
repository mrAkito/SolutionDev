CREATE TABLE IF NOT EXISTS Videos(
    id int AUTO_INCREMENT PRIMARY KEY,
    current_step_id int,
    next_step_id int,
    previous_step_id int,
    url longtext,
    FOREIGN KEY (current_step_id) REFERENCES Steps(id)
);