"use server";

export default async function addUser(formData: FormData) {
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
