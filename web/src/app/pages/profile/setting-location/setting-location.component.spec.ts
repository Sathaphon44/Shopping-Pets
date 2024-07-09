import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SettingLocationComponent } from './setting-location.component';

describe('SettingLocationComponent', () => {
  let component: SettingLocationComponent;
  let fixture: ComponentFixture<SettingLocationComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SettingLocationComponent]
    });
    fixture = TestBed.createComponent(SettingLocationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
