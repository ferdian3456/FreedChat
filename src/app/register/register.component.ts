import { Component, OnInit } from "@angular/core";
import { ApiService } from "../api.service";
import { Router } from "@angular/router";

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {
  data: any = {
    name: "",
    password: "",
  }

  constructor(
    public api: ApiService,
    public router: Router,
  ) {}

  register(){
    console.log("Register") 
  }
}
