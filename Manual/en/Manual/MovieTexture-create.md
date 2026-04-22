# Create a Movie Texture

To create a Movie Texture, place a video file in your project’s **Assets Folder**. Unity uses this video file in the same way as a regular [Texture](class-TextureImporter.md).

Unity imports video files using Apple QuickTime. On Windows, you need to install Quicktime to import a video file. Download Quicktime from [Apple Support Downloads](https://support.apple.com/downloads/quicktime). Unity supports the same file types as your QuickTime installation (usually **.mov**, **.mpg**, **.mpeg**, **.mp4**, **.avi**, **.asf**).

When you add a video file to your Project, Unity automatically imports it and converts it to **Ogg Theora** format. Once Unity has imported your Movie Texture, you can attach it to any **GameObject** or **Material** in the same way as a regular Texture.

## Playing the Movie

Your Movie Texture will not play automatically when the game begins running. You must use a short script to tell it when to play.

```
// this line of code will make the Movie Texture begin playing
((MovieTexture)GetComponent<Renderer>().material.mainTexture).Play();
```

Attach the following script to toggle Movie playback when the space bar is pressed:

```
public class PlayMovieOnSpace : MonoBehaviour {
    void Update () {
        if (Input.GetButtonDown ("Jump")) {
            
            Renderer r = GetComponent<Renderer>();
            MovieTexture movie = (MovieTexture)r.material.mainTexture;
            
            if (movie.isPlaying) {
                movie.Pause();
            }
            else {
                movie.Play();
            }
        }
    }
}
```

For more information about playing Movie Textures, see the [Movie Texture Script Reference](../ScriptReference/MovieTexture.md) page

### Movie Audio

When you import a Movie Texture, Unity also imports the accompanying audio track. This audio appears as an **AudioClip** child of the Movie Texture.

![The videos audio track appears as a child of the Movie Texture in the Project View](../uploads/Main/MovieTextureAudio.png)

The video’s audio track appears as a child of the Movie Texture in the Project View

To play this audio, the **Audio Clip** must be attached to a GameObject. Drag the Audio Clip from the Project View onto any GameObject in the **Scene** or Hierarchy View. Usually, this will be the same GameObject that is showing the Movie. Then use [AudioSource.Play()](../ScriptReference/AudioSource.Play.md) to make the the movie’s audio track play along with its video.