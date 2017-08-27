import { Component, OnInit } from '@angular/core';
import { Eleccion } from '../../models/eleccion';
import { EleccionService } from '../../services/eleccion.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-eleccion-new',
  templateUrl: './eleccion-new.component.html',
  styleUrls: ['./eleccion-new.component.css']
})
export class EleccionNewComponent implements OnInit {

  eleccion: Eleccion;
  display = false;
  constructor(private eleccionService: EleccionService, private location: Location) { }

  ngOnInit() {
    this.eleccion = new Eleccion();
  }

  guardar(eleccion: Eleccion): void {

    this.eleccionService.create(eleccion);
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