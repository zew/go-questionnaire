
SELECT 
    frage_kurz, 

    GROUP_CONCAT(
        DISTINCT antwort
        ORDER BY antwort
        SEPARATOR ','
    ) antw,

    count(*) anz


FROM `sonderfragen_ger` 
 WHERE `survey_id` = 202004
group by frage_kurz

ORDER BY `survey_id`  DESC, `frage_kurz` asc




