import { Component, OnInit } from '@angular/core';
import { Direccion } from '../../models/direccion';
import { DireccionService } from '../../services/direccion.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-direccion-edit',
  templateUrl: './direccion-edit.component.html',
  styleUrls: ['./direccion-edit.component.css']
})
export class DireccionEditComponent implements OnInit {

  direccion: Direccion = new Direccion();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private direccionService: DireccionService) {

  }

  actualizar(direccion: Direccion): void {
    this.direccionService.update(direccion).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.direccionService.getDireccion(params['id']))
      .subscribe(direccion => this.direccion = direccion);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}