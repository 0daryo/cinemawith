INSERT INTO `
cinemawith`.users
(
id,
uid,
pr,
name,
profile_url,
sex,
age,
nickname,
favorite,
enabled,created_at,updated_at
)
VALUES
    (1, '1111', 'よろしく', '山田太郎', 'http', 0, 21, "yama", "batman", true, 555555, 555555),
    (2, '2222', 'よろしく', '佐藤二郎', 'http', 0, 21, "yama", "batman", true, 555555, 555555),
    (3, '3333', 'よろしく', '田中三郎', 'http', 0, 21, "yama", "batman", true, 555555, 555555);


INSERT INTO `
cinemawith`.movies
(
id,
title,
description,
genre,
url,
enabled,created_at,updated_at
)
VALUES
    (1, 'アベンジャーズ', 'アイアンマン', "アクション", "http", true, 555555, 555555),
    (2, 'アイアンマン', 'アイアンマン', "アクション", "http", true, 555555, 555555),
    (3, 'ダークナイト', 'アイアンマン', "アクション", "http", true, 555555, 555555);


INSERT INTO `
cinemawith`.friends
(
id,
user_id,
friend_id,
enabled,created_at,updated_at
)
VALUES
    (1, 1, 2, true, 555555, 555555),
    (2, 2, 1, true, 555555, 555555);

INSERT INTO `
cinemawith`.users_movies
(
id,
user_id,
movie_id,
enabled,created_at,updated_at
)
VALUES
    (1, 1, 2, true, 555555, 555555),
    (2, 2, 3, true, 555555, 555555),
    (3, 3, 1, true, 555555, 555555);