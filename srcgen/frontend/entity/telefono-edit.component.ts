import { Component, OnInit } from '@angular/core';
import { Telefono } from '../../models/telefono';
import { TelefonoService } from '../../services/telefono.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-telefono-edit',
  templateUrl: './telefono-edit.component.html',
  styleUrls: ['./telefono-edit.component.css']
})
export class TelefonoEditComponent implements OnInit {

  telefono: Telefono = new Telefono();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private telefonoService: TelefonoService) {

  }

  actualizar(telefono: Telefono): void {
    this.telefonoService.update(telefono).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.telefonoService.getTelefono(params['id']))
      .subscribe(telefono => this.telefono = telefono);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}