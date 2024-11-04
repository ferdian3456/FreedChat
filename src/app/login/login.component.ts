import { Component, OnInit } from "@angular/core";
import { ApiService } from "../api.service";
import { Router } from "@angular/router";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  data: any = {
    name: "",
    password: "",
  }

  constructor(
    public api: ApiService,
    public router: Router,
  ) {}

  login(){
    console.log("Login")
  }
}
