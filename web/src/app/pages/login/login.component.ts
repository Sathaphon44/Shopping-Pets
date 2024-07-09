import { Component } from '@angular/core';
import { FormControl, FormGroup, NonNullableFormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/services/api/api.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent {
  
  constructor(
    private fb: NonNullableFormBuilder,
    private http: ApiService,
    private router: Router
  ) {}


  validateForm: FormGroup<{
    email: FormControl<string>;
    password: FormControl<string>;
    remember: FormControl<boolean>;
  }> = this.fb.group({
    email: ['', [Validators.required]],
    password: ['', [Validators.required]],
    remember: [true]
  });
  
  submitForm(): void {
    const email: string | undefined = this.validateForm.value.email;
    const password: string | undefined = this.validateForm.value.password;
    if (email !== undefined && password !== undefined) {
      this.http.LoginServices(email, password).subscribe((res: any) => {
        localStorage.setItem("token", res.token)
        this.router.navigate(["/"]).then(()=> {window.location.reload()})
      });
    };
  };






}
