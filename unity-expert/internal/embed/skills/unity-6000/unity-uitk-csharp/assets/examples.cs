using UnityEngine;
using UnityEngine.UIElements;

namespace MyGame.UI
{
    /// <summary>
    /// Example runtime UI controller for Unity UI Toolkit.
    /// Unity 6000.3 LTS compatible.
    /// </summary>
    public class ExampleUIController : MonoBehaviour
    {
        private UIDocument m_Document;
        private Label m_StatusLabel;
        private Button m_ConfirmButton;
        private Button m_CancelButton;
        private TextField m_NameInput;
        private Toggle m_EnabledToggle;

        private void OnEnable()
        {
            m_Document = GetComponent<UIDocument>();
            var root = m_Document.rootVisualElement;

            // Query elements
            m_StatusLabel = root.Q<Label>("status-label");
            m_ConfirmButton = root.Q<Button>("confirm-button");
            m_CancelButton = root.Q<Button>("cancel-button");
            m_NameInput = root.Q<TextField>("name-input");
            m_EnabledToggle = root.Q<Toggle>("enabled-toggle");

            // Register callbacks
            m_ConfirmButton.clicked += OnConfirm;
            m_CancelButton.clicked += OnCancel;
            m_NameInput.RegisterCallback<ChangeEvent<string>>(OnNameChanged);
            m_EnabledToggle.RegisterValueChangedCallback(OnToggleChanged);

            // Initial state
            m_StatusLabel.text = "Ready";
        }

        private void OnDisable()
        {
            if (m_Document == null) return;
            var root = m_Document.rootVisualElement;

            // Unregister callbacks
            m_ConfirmButton.clicked -= OnConfirm;
            m_CancelButton.clicked -= OnCancel;
            m_NameInput.UnregisterCallback<ChangeEvent<string>>(OnNameChanged);
            m_EnabledToggle.UnregisterValueChangedCallback(OnToggleChanged);
        }

        private void OnConfirm()
        {
            var name = m_NameInput.value;
            var enabled = m_EnabledToggle.value;
            m_StatusLabel.text = $"Confirmed: {name} (enabled: {enabled})";
            Debug.Log($"Confirm clicked: name={name}, enabled={enabled}");
        }

        private void OnCancel()
        {
            m_NameInput.value = "";
            m_EnabledToggle.SetValueWithoutNotify(false);
            m_StatusLabel.text = "Cancelled";
        }

        private void OnNameChanged(ChangeEvent<string> evt)
        {
            m_StatusLabel.text = $"Typing: {evt.newValue}";
        }

        private void OnToggleChanged(ChangeEvent<bool> evt)
        {
            m_StatusLabel.text = $"Enabled: {evt.newValue}";
        }
    }

    /// <summary>
    /// Example custom control using Unity 6 [UxmlElement] attribute system.
    /// </summary>
    [UxmlElement]
    public partial class HealthBar : VisualElement
    {
        public static readonly string ussClassName = "health-bar";

        private ProgressBar m_ProgressBar;
        private Label m_Label;
        private float m_CurrentValue;
        private float m_MaxValue = 100f;

        [UxmlAttribute("current-value")]
        public float CurrentValue
        {
            get => m_CurrentValue;
            set
            {
                m_CurrentValue = Mathf.Clamp(value, 0, m_MaxValue);
                UpdateDisplay();
            }
        }

        [UxmlAttribute("max-value")]
        public float MaxValue
        {
            get => m_MaxValue;
            set
            {
                m_MaxValue = Mathf.Max(value, 0.01f);
                UpdateDisplay();
            }
        }

        public HealthBar()
        {
            AddToClassList(ussClassName);

            m_ProgressBar = new ProgressBar
            {
                style = { flexGrow = 1 }
            };
            Add(m_ProgressBar);

            m_Label = new Label
            {
                style = { minWidth = 60, unityTextAlign = TextAnchor.MiddleCenter }
            };
            Add(m_Label);
        }

        private void UpdateDisplay()
        {
            m_ProgressBar.value = m_CurrentValue;
            m_ProgressBar.highValue = m_MaxValue;
            m_Label.text = $"{m_CurrentValue:F0}/{m_MaxValue:F0}";
        }
    }

    /// <summary>
    /// Example editor window pattern.
    /// </summary>
#if UNITY_EDITOR
    using UnityEditor;

    public class ExampleEditorWindow : EditorWindow
    {
        [SerializeField] private VisualTreeAsset m_VisualTreeAsset;

        [MenuItem("Tools/Example Window")]
        public static void ShowWindow()
        {
            var wnd = GetWindow<ExampleEditorWindow>();
            wnd.titleContent = new GUIContent("Example Window");
            wnd.minSize = new Vector2(400, 300);
        }

        public void CreateGUI()
        {
            VisualElement root = rootVisualElement;

            if (m_VisualTreeAsset != null)
            {
                m_VisualTreeAsset.CloneTree(root);
            }

            var button = root.Q<Button>("confirm-button");
            if (button != null)
            {
                button.clicked += () => Debug.Log("Editor button clicked!");
            }
        }
    }
#endif
}
