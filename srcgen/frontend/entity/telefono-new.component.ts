import { Component, OnInit } from '@angular/core';
import { Telefono } from '../../models/telefono';
import { TelefonoService } from '../../services/telefono.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-telefono-new',
  templateUrl: './telefono-new.component.html',
  styleUrls: ['./telefono-new.component.css']
})
export class TelefonoNewComponent implements OnInit {

  telefono: Telefono;
  display = false;
  constructor(private telefonoService: TelefonoService, private location: Location) { }

  ngOnInit() {
    this.telefono = new Telefono();
  }

  guardar(telefono: Telefono): void {

    this.telefonoService.create(telefono);
    this.display = true;

  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}