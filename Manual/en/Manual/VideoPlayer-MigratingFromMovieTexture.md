# Migrating from MovieTexture to VideoPlayer

If you have projects that use the legacy [MovieTexture](class-MovieTexture.md) component to download and play movies, you can update them to use the newer [VideoPlayer](class-VideoPlayer.md) component.

To help you migrate from MovieTexture to VideoPlayer, this page provides examples of how to download and play back a movie using each component.

## Playing a Movie

### MovieTexture

```
using UnityEngine;

public class PlayMovieMT : MonoBehaviour
{
    public AudioClip movieAudioClip;
    public MovieTexture movieTexture;

    void Start()
    {
        var audioSource = gameObject.AddComponent<AudioSource>();
        audioSource.clip = movieAudioClip;
    }

    void Update()
    {
        if (Input.GetButtonDown("Jump"))
        {
            var audioSource = GetComponent<AudioSource>();
            GetComponent<Renderer>().material.mainTexture = movieTexture;

            if (movieTexture.isPlaying)
            {
                movieTexture.Pause();
                audioSource.Pause();
            }
            else
            {
                movieTexture.Play();
                audioSource.Play();
            }
        }
    }
}
```

### VideoPlayer

```
using UnityEngine;

public class PlayMovieVP : MonoBehaviour
{
    public UnityEngine.Video.VideoClip videoClip;

    void Start()
    {
        var videoPlayer = gameObject.AddComponent<UnityEngine.Video.VideoPlayer>();
        var audioSource = gameObject.AddComponent<AudioSource>();

        videoPlayer.playOnAwake = false;
        videoPlayer.clip = videoClip;
        videoPlayer.renderMode = UnityEngine.Video.VideoRenderMode.MaterialOverride;
        videoPlayer.targetMaterialRenderer = GetComponent<Renderer>();
        videoPlayer.targetMaterialProperty = "_MainTex";
        videoPlayer.audioOutputMode = UnityEngine.Video.VideoAudioOutputMode.AudioSource;
        videoPlayer.SetTargetAudioSource(0, audioSource);
    }

    void Update()
    {
        if (Input.GetButtonDown("Jump"))
        {
            var vp = GetComponent<UnityEngine.Video.VideoPlayer>();

            if (vp.isPlaying)
            {
                vp.Pause();
            }
            else
            {
                vp.Play();
            }
        }
    }
}
```

## Downloading a Movie

### MovieTexture

```
using UnityEngine;

public class DownloadMovieMT : MonoBehaviour
{
    void Start()
    {
        StartCoroutine(GetMovieTexture());
    }

    IEnumerator GetMovieTexture()
    {
        using (var uwr = UnityWebRequestMultimedia.GetMovieTexture("https://myserver.com/mymovie.ogv"))
        {
            yield return uwr.SendWebRequest();
            if (uwr.isNetworkError || uwr.isHttpError)
            {
                Debug.LogError(uwr.error);
                yield break;
            }

            MovieTexture movie = DownloadHandlerMovieTexture.GetContent(uwr);

            GetComponent<Renderer>().material.mainTexture = movie;
            movie.loop = true;
            movie.Play();
        }
    }
}
```

### VideoPlayer

```
using UnityEngine;

public class DownloadMovieVP : MonoBehaviour
{
    void Start()
    {
        var vp = gameObject.AddComponent<UnityEngine.Video.VideoPlayer>();
        vp.url = "https://myserver.com/mymovie.mp4";

        vp.isLooping = true;
        vp.renderMode = UnityEngine.Video.VideoRenderMode.MaterialOverride;
        vp.targetMaterialRenderer = GetComponent<Renderer>();
        vp.targetMaterialProperty = "_MainTex";

        vp.Play();
    }
}
```