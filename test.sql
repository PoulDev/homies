INSERT INTO avatars (bg_color, face_color, face_x, face_y, left_eye_x, left_eye_y, right_eye_x, right_eye_y, bezier)
VALUES (
    '000000',
    '000000',
    0.0,
    0.0,
    0.0,
    0.0,
    0.0,
    0.0,
    '2 2 5 1 6 0'
);

INSERT INTO users (name, email, pwd_hash, pwd_salt, avatar)
VALUES (
    'Traba',
    'john.doe1337@duck.com',
    'c69734f24b0781ebae4adb7a137c07705d6d9f6f0a68f20973f2a5d834cd55ae',
    'a98a31fd4804c7cb5f0c9b64ae6c4ba8',
    1
)
