SELECT random(1);
SELECT position();

-- stderr
-- # package querytest
-- query.sql:1:8: function random(unknown) does not exist
-- query.sql:2:8: function position() does not exist
