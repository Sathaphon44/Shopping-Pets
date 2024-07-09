import { Component } from '@angular/core';
import { FormControl, FormGroup, NonNullableFormBuilder, Validators } from '@angular/forms';
import { ApiService } from 'src/app/services/api/api.service';


@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {

  validateForm: FormGroup<{
    username: FormControl<string>;
    email: FormControl<string>;
    createPassword: FormControl<string>;
    confirmPassword: FormControl<string>;
  }> = this.fb.group({
    username: ['', [Validators.required]],
    email: ['', [Validators.required]],
    createPassword: ['', [Validators.required]],
    confirmPassword: ['', [Validators.required]],
  });

  constructor(
    private fb: NonNullableFormBuilder,
    private api: ApiService
  ) { }


  submitForm(): void {
    const { username, email, createPassword, confirmPassword } = this.validateForm.value;
    if (username && email && createPassword && confirmPassword) {
      this.api.RegisterService(username, email, confirmPassword).subscribe((res: any) => {
        console.log(res)
      })
    }
  }

}
