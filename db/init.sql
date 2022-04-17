INSERT INTO users (nickname, email, password) VALUES
    --hashlib.sha256(b'aye1488228').hexdigest()
    --('modnaya_shalava', 'modnaya.shalava@gmail.com', '105533113bbec8d922cf934db349762fcf251c3bac486968d2112b41a8754edc'),
    ('modnaya_shalava', 'modnaya.shalava@gmail.com', '68687364616875686b75105533113bbec8d922cf934db349762fcf251c3bac486968d2112b41a8754edc'),
    --qwerty123
    --('big_boy', 'bigboy@gmail.com', 'daaad6e5604e8e17bd9f108d91e26afe6281dac8fda0091040a7a6d7bd9b43b5');
    ('big_boy', 'bigboy@gmail.com', '68687364616875686b75daaad6e5604e8e17bd9f108d91e26afe6281dac8fda0091040a7a6d7bd9b43b5');

INSERT INTO clothes (photo_id, owner_id, class, brand, color) VALUES 
    ('dsagfddgf', 1, 'jacket', 'Patagonia', 0),
    ('dsadsafaf', 1, 'pants', 'Adidas', 1);

INSERT INTO looks (photo_id, owner_id, description, season, temperature_range, purpose, priority) VALUES
    (111111111, 1, 'Perfect look for wonderful day', 'fall/spring', int4range(5, 15), 'casual', 1);

INSERT INTO looks_clothes (look_id, cloth_id) VALUES
    (1, 1),
    (1, 2);
    