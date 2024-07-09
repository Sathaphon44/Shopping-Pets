import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { en_US } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import en from '@angular/common/locales/en';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { IconsProviderModule } from './icons-provider.module';

// components
import { LoginComponent } from './pages/login/login.component';
import { NavbarComponent } from './app-components/navbar/navbar.component';
import { RegisterComponent } from './pages/register/register.component';
import { FooterComponent } from './app-components/footer/footer.component';
import { ShopComponent } from './pages/shop/shop.component';

// ant-design
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { NzCardModule } from 'ng-zorro-antd/card';
import { CurrentUserService } from './services/current-user/current-user.service';
import { CartComponent } from './pages/cart/cart.component';
import { AuthGuard } from './services/auth-guard/auth-guard.service';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzBadgeModule } from 'ng-zorro-antd/badge';
import { DataSharingService } from './services/data-sharing/data-sharing.service';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzAvatarModule } from 'ng-zorro-antd/avatar';
import { ProfileComponent } from './pages/profile/profile.component';
import { SettingProfileComponent } from './pages/profile/setting-profile/setting-profile.component';
import { SettingLocationComponent } from './pages/profile/setting-location/setting-location.component';

registerLocaleData(en);

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    NavbarComponent,
    RegisterComponent,
    FooterComponent,
    ShopComponent,
    CartComponent,
    ProfileComponent,
    SettingProfileComponent,
    SettingLocationComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    IconsProviderModule,
    NzLayoutModule,
    NzMenuModule,
    FormsModule,
    NzFormModule,
    ReactiveFormsModule,
    NzInputModule,
    NzButtonModule,
    NzCheckboxModule,
    NzGridModule,
    NzCardModule,
    HttpClientModule,
    NzModalModule,
    NzBadgeModule,
    NzIconModule,
    NzAvatarModule
  ],
  providers: [
    { provide: NZ_I18N, useValue: en_US },
    CurrentUserService,
    AuthGuard,
    DataSharingService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
