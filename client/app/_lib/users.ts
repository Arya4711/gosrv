"use server";
import { User } from "./models";

export async function getUser(formData: FormData) {
  const username = formData.get("username") as string;
  const response = await fetch(`http://localhost:5050/users/${username}`, {
    method: "GET",
  });

  return (await response.json()) as User;
}

export async function getUsers() {
  const response = await fetch("http://localhost:5050/users", {
    method: "GET",
  });

  return (await response.json()) as User[];
}

export async function deleteUser(formData: FormData) {
  const username = formData.get("username");

  await fetch("http://localhost:5050/users", {
    method: "DELETE",
    body: JSON.stringify({ username }),
  });
}

export async function addUser(formData: FormData) {
  const username = formData.get("username") as string;
  const password = formData.get("password") as string;

  await fetch("http://localhost:5050/users", {
    method: "POST",
    body: JSON.stringify({
      username,
      password,
    }),
  });
}
