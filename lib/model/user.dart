library user;

class User {
  String displayName;
  String email;
  String password;

  User(this.displayName, this.email, this.password);

  bool operator==(other) {
    if (other is! User) {
      return false;
    }
    User user = other as User;
    return (displayName == user.displayName &&
            email == user.email &&
            password == user.password);
  }

  Map<String, dynamic> toJson() {
    var map = new Map<String, dynamic>();
    map['display_name'] = displayName;
    map['email']= email;
    map['password'] = password;
    return map;
  }
}