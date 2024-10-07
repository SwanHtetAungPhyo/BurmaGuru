create table if not exists User ( 
    Id int primary key autoincrement,
    user_name text not null, 
    email text not null, 
    p_assword text not null, 
    profilePicture text not null, 
    token text
    isVerified boolean default false, 
)