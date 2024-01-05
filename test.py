import requests
import json

base_url = "http://localhost:8000"

# GET /users
def test_get_users():
    response = requests.get(f"{base_url}/users")
    assert response.status_code == 200
    users = response.json()
    assert isinstance(users, list)

# GET /users/{id}
# check id
def test_get_user():
    response = requests.get(f"{base_url}/users/6")
    assert response.status_code == 200
    user = response.json()
    assert isinstance(user, dict)
    assert user["id"] == 6

# POST /users
def test_create_user():
    new_user = {"username": "new_user", "email": "new_user@example.com"}
    response = requests.post(f"{base_url}/users", json=new_user)
    assert response.status_code == 200
    user = response.json()
    assert isinstance(user, dict)
    assert "id" in user
    assert user["username"] == new_user["username"]
    assert user["email"] == new_user["email"]

# PUT /users/{id}
def test_update_user():
    updated_user = {"username": "updated_user", "email": "updated_user@example.com"}
    response = requests.put(f"{base_url}/users/1", json=updated_user)
    assert response.status_code == 200
    user = response.json()
    assert isinstance(user, dict)
    assert user["id"] == 0
    assert user["username"] == updated_user["username"]
    assert user["email"] == updated_user["email"]

# DELETE /users/{id}
def test_delete_user():
    response = requests.delete(f"{base_url}/users/1")
    assert response.status_code == 204
    # Check if the user has been deleted by trying to retrieve it again
    response = requests.get(f"{base_url}/users/1")
    assert response.status_code == 404

if __name__ == "__main__":
    test_get_users()
    test_get_user()
    test_create_user()
    test_update_user()
    test_delete_user()
    print("All tests passed!")
