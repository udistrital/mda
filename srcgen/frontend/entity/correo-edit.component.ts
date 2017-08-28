import { Component, OnInit } from '@angular/core';
import { Correo } from '../../models/correo';
import { CorreoService } from '../../services/correo.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-correo-edit',
  templateUrl: './correo-edit.component.html',
  styleUrls: ['./correo-edit.component.css']
})
export class CorreoEditComponent implements OnInit {

  correo: Correo = new Correo();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private correoService: CorreoService) {

  }

  actualizar(correo: Correo): void {
    this.correoService.update(correo).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.correoService.getCorreo(params['id']))
      .subscribe(correo => this.correo = correo);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}