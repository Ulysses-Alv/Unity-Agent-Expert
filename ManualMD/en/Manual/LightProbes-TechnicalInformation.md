# Light Probe data format

The lighting information in the **light probes** are encoded as Spherical Harmonics basis functions. We use third order polynomials, also known as L2 Spherical Harmonics. These are stored using 27 floating point values, 9 for each color channel.

The **Enlighten** Realtime **Global Illumination** implementation in Unity uses a different spherical harmonics approach than Geomerics, the feature’s original developer, namely the notation and reconstruction method from Peter-Pike Sloan’s paper, [Stupid Spherical Harmonics (SH) Tricks](http://www.ppsloan.org/publications/StupidSH36.pdf). Geomerics’ original approach was based on the notation and reconstruction method from Ramamoorthi/Hanrahan’s paper, [An Efﬁcient Representation for Irradiance Environment Maps](http://cseweb.ucsd.edu/~ravir/papers/envmap/envmap.pdf).

The **shader** code for reconstruction is found in UnityCG.cginc and is using the method from Appendix A10 Shader/CPU code for Irradiance Environment Maps from Peter-Pikes paper.

The data is internally ordered like this:

```
                        [L00:  DC]

            [L1-1:  y] [L10:   z] [L11:   x]

  [L2-2: xy] [L2-1: yz] [L20:  zz] [L21:  xz]  [L22:  xx - yy]
```

The 9 coefficients for R, G and B are ordered like this:

```
L00, L1-1,  L10,  L11, L2-2, L2-1,  L20,  L21,  L22, // red channel

L00, L1-1,  L10,  L11, L2-2, L2-1,  L20,  L21,  L22, // blue channel

L00, L1-1,  L10,  L11, L2-2, L2-1,  L20,  L21,  L22  // green channel
```

For more “under-the-hood” information about Unity’s light probe system, you can read Robert Cupisz’s talk from GDC 2012, [Light Probe Interpolation Using Tetrahedral Tessellations”, GDC 2012](http://gdcvault.com/play/1015312/Light-Probe-Interpolation-Using-Tetrahedral)