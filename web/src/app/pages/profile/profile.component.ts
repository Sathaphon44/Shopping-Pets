import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {

  chooseMenu: any = {
    Setting: true,
    Location: true
  };

  constructor(
    private router: Router,
    private AcRouter: ActivatedRoute
  ){}

  handleClickMenu(intent: string) {
    this.chooseMenu.Setting = (intent == "setting")
    this.chooseMenu.Location = (intent == "location")
    this.router.navigate(['profile', intent])
  }

  checkPathCurrent() {
    this.AcRouter.params.subscribe((params: any) => {
      this.chooseMenu.Setting = (params.intent == "setting")
      this.chooseMenu.Location = (params.intent == "location")
    })
  }


  ngOnInit(): void {
    //Called after the constructor, initializing input properties, and the first call to ngOnChanges.
    //Add 'implements OnInit' to the class.
    this.checkPathCurrent()
  }
}
