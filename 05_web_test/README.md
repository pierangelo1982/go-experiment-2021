SELECT c.id, c.firstName, ci.lastName, c.carNumber, c.teamId json_agg(json_build_object('name', country)) AS team
    FROM Driver AS c 
    LEFT JOIN Team AS t 
    ON t.id=c.teamId 
    GROUP BY c.id;