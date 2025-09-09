"use client";

import { Button } from "@/components/ui/button";
import { loginUser } from "@/helper/auth.helper";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { toast } from "sonner";

export default function Login() {
    const router = useRouter();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const response = await loginUser(username, password);
      console.log("Login success:", response);

      // toast("Login Succes, Welcome Back !", username);
      toast.success("Login Success", {
        description: `Welcome back, ${username}!`,
      });

      // Example: store token if API returns one
      // localStorage.setItem("token", response.token);

      router.push("/");
    } catch (err: any) {
      console.error(err);
      setError(err || "Login failed");
    } finally {
      setLoading(false);
    }
  };
  return (
    <div className="flex justify-center items-center w-full bg-gray-100">
      <div className=" w-fit h-fit bg-white rounded-xl shadow-lg px-32 py-24">
        <div className="text-center mb-6">
          <p className="text-3xl ">Welcome Back</p>
          <p className="text-gray-500 mt-2">
            Dont have an account yet? <Link href="/register">Sign Up</Link>
          </p>
        </div>
        <form onSubmit={handleSubmit} className="max-w-sm mx-auto">
          <div className="mb-5">
            <label
              htmlFor="username"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Username
            </label>
            <input
              type="text"
              id="text"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Insert your username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="mb-5">
            <label
              htmlFor="password"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Password
            </label>
            <input
              type="password"
              id="password"
              placeholder="Insert your password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              required
            />
          </div>
          {/* <button
            type="submit"
            className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Submit
          </button> */}
          {error && <p className="text-red-500 text-sm">{error}</p>}
          <Button className="w-full mt-2" disabled={loading}>
            {loading ? "Logging in..." : "Submit"}
          </Button>
        </form>
      </div>
    </div>
  );
}
